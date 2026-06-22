// 在整站「多页刷新」的前提下，记忆并还原左侧目录（book-menu）的两类状态：
//   1) 滚动位置：避免每次翻页后目录跳回顶部；
//   2) 展开/收起状态：让你手动展开的多个分组在翻页后依旧保持展开（去掉“只保留当前页路径”的收起效果）。
(function () {
  var SCROLL_KEY = "menu.scrollTop";
  var TOGGLE_KEY = "menu.toggles";

  // 真正可滚动的是 .book-menu-content（position:fixed; overflow-y:auto），不是它内部的 <nav>。
  function getMenu() {
    return (
      document.querySelector("aside.book-menu .book-menu-content") ||
      document.querySelector("aside.book-menu nav")
    );
  }

  // ---------- 展开/收起状态 ----------
  function loadToggles() {
    try {
      return JSON.parse(localStorage.getItem(TOGGLE_KEY) || "{}") || {};
    } catch (e) {
      return {};
    }
  }

  function restoreToggles() {
    var saved = loadToggles();
    var toggles = document.querySelectorAll("aside.book-menu .toggle");
    for (var i = 0; i < toggles.length; i++) {
      var cb = toggles[i];
      // checkbox 的 id 形如 section-<md5>，跨页面稳定，可作为 key。
      if (cb.id && Object.prototype.hasOwnProperty.call(saved, cb.id)) {
        cb.checked = saved[cb.id];
      }
    }
    // 始终保证当前页（.active）所在的所有可折叠祖先保持展开，避免看不到当前条目。
    var active = document.querySelector("aside.book-menu a.active");
    if (active) {
      var node = active.parentElement;
      while (node && node.closest && node.closest("aside.book-menu")) {
        if (node.tagName === "LI") {
          var t = node.querySelector(":scope > .toggle");
          if (t) {
            t.checked = true;
          }
        }
        node = node.parentElement;
      }
    }
  }

  function bindToggles() {
    var toggles = document.querySelectorAll("aside.book-menu .toggle");
    for (var i = 0; i < toggles.length; i++) {
      (function (cb) {
        cb.addEventListener("change", function () {
          var saved = loadToggles();
          saved[cb.id] = cb.checked;
          localStorage.setItem(TOGGLE_KEY, JSON.stringify(saved));
        });
      })(toggles[i]);
    }
  }

  // ---------- 滚动位置 ----------
  function restoreScroll() {
    var menu = getMenu();
    if (!menu) {
      return;
    }
    var saved = localStorage.getItem(SCROLL_KEY);
    if (saved !== null) {
      menu.scrollTop = parseInt(saved, 10) || 0;
    }
  }

  var scrollBound = false;
  function bindScroll() {
    if (scrollBound) {
      return;
    }
    var menu = getMenu();
    if (!menu) {
      return;
    }
    scrollBound = true;
    menu.addEventListener(
      "scroll",
      function () {
        localStorage.setItem(SCROLL_KEY, menu.scrollTop);
      },
      { passive: true }
    );
  }

  // 关键：闪烁的根因是“在右侧内容（window load）加载完之前，目录一直没能被还原”。
  // 这里把滚动位置「持续钉住」——每一帧都把它设回记忆值，直到 load 之后再停。
  // 但是：一旦用户主动操作目录（滚轮 / 触摸 / 拖动滚动条 / 按键），必须立刻停止钉位，
  // 否则会和用户的滚动“打架”，表现为目录抖动、被拽回记忆位置。
  // 程序化设置 scrollTop 只会触发 scroll 事件、不会触发 wheel/touch/pointer/key，
  // 所以用后者判断“用户意图”，可与钉位自身的滚动区分开。
  var savedScroll = parseInt(localStorage.getItem(SCROLL_KEY) || "0", 10) || 0;
  var stopPinAt = null; // load 之后再多钉一小段时间，然后停止
  var frames = 0;
  var userInteracted = false;

  function stopPinning() {
    if (userInteracted) {
      return;
    }
    userInteracted = true;
    bindScroll(); // 立刻开始记录用户的真实滚动
  }

  function bindInteraction() {
    var menu = getMenu();
    var target = menu || document;
    ["wheel", "touchstart", "pointerdown", "keydown"].forEach(function (ev) {
      target.addEventListener(ev, stopPinning, { passive: true });
    });
  }

  function pin() {
    if (userInteracted) {
      return; // 用户已开始操作，交还控制权
    }
    var menu = getMenu();
    if (menu && savedScroll > 0 && menu.scrollTop !== savedScroll) {
      menu.scrollTop = savedScroll;
    }
    frames++;
    // 停止条件：load 之后再钉 ~20 帧；或兜底最多 ~600 帧，避免极端情况下死循环。
    if (stopPinAt !== null && frames > stopPinAt) {
      bindScroll();
      return;
    }
    if (frames < 600) {
      requestAnimationFrame(pin);
    } else {
      bindScroll();
    }
  }

  window.addEventListener("load", function () {
    if (stopPinAt === null) {
      stopPinAt = frames + 20;
    }
    if (!userInteracted) {
      restoreScroll();
    }
  });

  // 解析到此立即还原一次（展开状态需在滚动之前还原，因为它会改变目录总高度），并启动钉位。
  restoreToggles();
  restoreScroll();
  bindInteraction();
  requestAnimationFrame(pin);

  // DOM 就绪后再还原展开状态并绑定其保存（滚动的保存在钉位结束后才绑定），
  // 并补绑一次交互监听，防止内联执行时菜单元素尚未就绪。
  document.addEventListener("DOMContentLoaded", function () {
    restoreToggles();
    bindToggles();
    bindInteraction();
  });
})();
