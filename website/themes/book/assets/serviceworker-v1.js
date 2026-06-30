'use strict';
(function () {
  'use strict';
    /**
    * Service Worker Toolbox caching
    */
    // 缓存版本号。每次缓存策略变更都要 +1：activate 时会删除所有不在 currentCaches 里的旧缓存，
    // 让老客户端拿到新 SW 后立刻丢弃旧版本资源，避免一直被钉在旧页面上（曾导致 search.js 失焦修复进不来）。
    var cacheVersion = '-toolbox-v2';
    var dynamicVendorCacheName = 'dynamic-vendor' + cacheVersion;
    var staticVendorCacheName = 'static-vendor' + cacheVersion;
    var staticAssetsCacheName = 'static-assets' + cacheVersion;
    var contentCacheName = 'content' + cacheVersion;
    var currentCaches = [dynamicVendorCacheName, staticVendorCacheName, staticAssetsCacheName, contentCacheName];
    var maxEntries = 50;
    self.importScripts("/leetcode/js/sw-toolbox.js");
    self.toolbox.options.debug = false;

    // 同源请求统一入口（必须第一个注册，sw-toolbox 取第一条匹配的路由）：
    //   - 页面导航 / text/html  → networkFirst：始终优先取最新 HTML，从而引用到最新指纹的 JS/CSS，
    //     不会再把用户钉死在旧版本（这正是“失焦修复已部署但用户仍复现”的根因——旧 HTML 走 cacheFirst）。
    //     断网时回退缓存，PWA 离线能力保留。
    //   - 其余静态资源（文件名带内容指纹，内容变 URL 必变）→ cacheFirst，安全且最快。
    // networkTimeoutSeconds：networkFirst 等网络超过该秒数仍无响应，就回退缓存，避免“网络慢但没断”时干等白屏；
    //   仅对 networkFirst 生效，对 cacheFirst 分支无影响。
    self.toolbox.router.get('/(.*)', function (request, values, options) {
        var accept = request.headers.get('accept') || '';
        if (request.mode === 'navigate' || accept.indexOf('text/html') !== -1) {
            return self.toolbox.networkFirst(request, values, options);
        }
        return self.toolbox.cacheFirst(request, values, options);
    }, {
        networkTimeoutSeconds: 3,
        cache: {
          name: staticAssetsCacheName,
          maxEntries: maxEntries
        }
    });
    // cache dynamic vendor assets, things which have no other update mechanism like filename change/version hash
    self.toolbox.router.get('/css', self.toolbox.fastest, {
        origin: /fonts\.googleapis\.com/,
            cache: {
              name: dynamicVendorCacheName,
              maxEntries: maxEntries
            }
    });
    // Do not cache disqus
    self.toolbox.router.get('/(.*)', self.toolbox.networkOnly, {
        origin: /disqus\.com/
    });
    self.toolbox.router.get('/(.*)', self.toolbox.networkOnly, {
        origin: /disquscdn\.com/
    });
    // Cache all static vendor assets, e.g. fonts whose version is bind to the according url
    self.toolbox.router.get('/(.*)', self.toolbox.cacheFirst, {
        origin: /(fonts\.gstatic\.com|www\.google-analytics\.com|www\.googletagmanager\.com|cdnjs\.cloudflare\.com)/,
        cache: {
          name: staticVendorCacheName,
          maxEntries: maxEntries
        }
    });
    self.toolbox.router.get('/(.*)', self.toolbox.cacheFirst, {
        origin: /(img\.halfrost\.com)/,
        cache: {
          name: staticAssetsCacheName,
          maxEntries: maxEntries
        }
    });
    // immediately activate this serviceworker
    self.addEventListener('install', function (event) {
        return event.waitUntil(self.skipWaiting());
    });
    // 激活时删除旧版本（v1 等）遗留的缓存，并立即接管所有页面
    self.addEventListener('activate', function (event) {
        return event.waitUntil(
            self.caches.keys().then(function (names) {
                return Promise.all(names.filter(function (name) {
                    return currentCaches.indexOf(name) === -1;
                }).map(function (name) {
                    return self.caches.delete(name);
                }));
            }).then(function () {
                return self.clients.claim();
            })
        );
    });
})();
