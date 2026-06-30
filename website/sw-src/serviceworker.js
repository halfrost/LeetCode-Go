/**
 * Service Worker —— Workbox 7 版本（替代已停维护的 sw-toolbox）。
 *
 * 构建方式（injectManifest）：
 *   `hugo` 构建出 public/ 后，运行 `npm run sw`（= workbox injectManifest workbox-config.cjs），
 *   它会把 public/ 里 app shell 资源的「文件 + 内容版本号」注入下面那个 __WB_MANIFEST 占位符，
 *   生成最终的 public/serviceworker-v1.min.js。
 *
 * 运行时缓存策略与迁移前（sw-toolbox 版）保持等价：
 *   - 页面导航(HTML) → NetworkFirst，3s 超时回退缓存（避免“网络慢但没断”干等）
 *   - 同源静态资源(带内容指纹) → CacheFirst
 *   - Google Fonts 的 CSS → StaleWhileRevalidate（对应旧的 fastest）
 *   - gstatic / google-analytics / googletagmanager / cdnjs → CacheFirst（含不透明响应）
 *   - img.halfrost.com → CacheFirst
 * 额外新增：用 precache manifest 预缓存 app shell（CSS/JS/字体/图标），实现真正的离线。
 */

importScripts('/leetcode/js/workbox-v7.4.1/workbox-sw.js');

// 用自托管的 Workbox 运行时（copyLibraries 拷到 /leetcode/js/workbox-v7.4.1/），不依赖第三方 CDN。
workbox.setConfig({
  modulePathPrefix: '/leetcode/js/workbox-v7.4.1/',
  debug: false,
});

const { precaching, routing, strategies, expiration, cacheableResponse, core } = workbox;

core.setCacheNameDetails({ prefix: 'leetcode', suffix: 'v1' });

// 新 SW 一就绪就接管，并立即控制所有页面（与旧版 skipWaiting + clients.claim 等价）。
self.skipWaiting();
core.clientsClaim();

// ─────────────── 预缓存（app shell） ───────────────
// 下面的 __WB_MANIFEST 占位符由 workbox injectManifest 在构建时替换为带版本号的资源清单；
// 若未跑注入步骤则为 undefined，用 || [] 兜底，保证 SW 仍能正常工作（仅少了预缓存）。
precaching.precacheAndRoute(self.__WB_MANIFEST || []);
// 删除上一版 precache 留下的过期缓存。
precaching.cleanupOutdatedCaches();

// ─────────────── 运行时路由 ───────────────

// 页面导航：NetworkFirst + 3s 超时。在线取最新 HTML（从而引用到最新指纹资源，不会被钉在旧版），
// 慢/断网回退缓存。这正是 NavigationRoute 的用法。
routing.registerRoute(
  new routing.NavigationRoute(
    new strategies.NetworkFirst({
      cacheName: 'leetcode-pages',
      networkTimeoutSeconds: 3,
      plugins: [new expiration.ExpirationPlugin({ maxEntries: 50 })],
    })
  )
);

// 同源静态资源（未被预缓存命中的部分，如体积大的 search-data、内容图片等）→ CacheFirst。
// 文件名带内容指纹，内容变 URL 必变，cacheFirst 安全且最快。
routing.registerRoute(
  ({ request, url }) =>
    url.origin === self.location.origin && request.mode !== 'navigate',
  new strategies.CacheFirst({
    cacheName: 'leetcode-static',
    plugins: [new expiration.ExpirationPlugin({ maxEntries: 50 })],
  })
);

// Google Fonts 的 CSS（无指纹更新机制）→ StaleWhileRevalidate（对应旧版 fastest）。
routing.registerRoute(
  ({ url }) => url.origin === 'https://fonts.googleapis.com',
  new strategies.StaleWhileRevalidate({ cacheName: 'leetcode-google-fonts-css' })
);

// 第三方静态资源（字体/统计/CDN，URL 绑定版本）→ CacheFirst；允许缓存不透明响应(status 0)。
routing.registerRoute(
  ({ url }) =>
    /(^|\.)(fonts\.gstatic\.com|www\.google-analytics\.com|www\.googletagmanager\.com|cdnjs\.cloudflare\.com)$/.test(
      url.hostname
    ),
  new strategies.CacheFirst({
    cacheName: 'leetcode-static-vendor',
    plugins: [
      new cacheableResponse.CacheableResponsePlugin({ statuses: [0, 200] }),
      new expiration.ExpirationPlugin({ maxEntries: 50 }),
    ],
  })
);

// 图床 img.halfrost.com → CacheFirst。
routing.registerRoute(
  ({ url }) => url.hostname === 'img.halfrost.com',
  new strategies.CacheFirst({
    cacheName: 'leetcode-img',
    plugins: [
      new cacheableResponse.CacheableResponsePlugin({ statuses: [0, 200] }),
      new expiration.ExpirationPlugin({ maxEntries: 50 }),
    ],
  })
);

// 清掉上一代 sw-toolbox 留下的缓存（名字含 -toolbox-），让老用户彻底切换干净。
self.addEventListener('activate', (event) => {
  event.waitUntil(
    caches.keys().then((keys) =>
      Promise.all(
        keys
          .filter((key) => key.indexOf('-toolbox-') !== -1)
          .map((key) => caches.delete(key))
      )
    )
  );
});
