'use strict';
(function () {
  'use strict';
    /**
    * Service Worker Toolbox caching
    */
    var cacheVersion = '-toolbox-v1';
    var dynamicVendorCacheName = 'dynamic-vendor' + cacheVersion;
    var staticVendorCacheName = 'static-vendor' + cacheVersion;
    var staticAssetsCacheName = 'static-assets' + cacheVersion;
    var contentCacheName = 'content' + cacheVersion;
    var maxEntries = 50;
    self.importScripts("/leetcode/js/sw-toolbox.js");
    self.toolbox.options.debug = false;

    // Cache own static assets
    self.toolbox.router.get('/(.*)', self.toolbox.cacheFirst, {
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
    self.toolbox.router.get('/content/(.*)', self.toolbox.fastest, {
        cache: {
          name: contentCacheName,
          maxEntries: maxEntries
        }
    });
    self.toolbox.router.get('/*', function (request, values, options) {
        if (!request.url.match(/(\/ghost\/|\/page\/)/) && request.headers.get('accept').includes('text/html')) {
            return self.toolbox.fastest(request, values, options);
        } else {
          // DevTools opening will trigger these o-i-c requests, which this SW can't handle.
          // There's probaly more going on here, but I'd rather just ignore this problem. :)
          // https://github.com/paulirish/caltrainschedule.io/issues/49
          if (request.cache === 'only-if-cached' && request.mode !== 'same-origin'){
            return;
          } else {
            return self.toolbox.networkOnly(request, values, options);
          }
        }
        }, {
        cache: {
            name: contentCacheName,
            maxEntries: maxEntries
        }
    });
    // immediately activate this serviceworker
    self.addEventListener('install', function (event) {
        return event.waitUntil(self.skipWaiting());
    });
    self.addEventListener('activate', function (event) {
        return event.waitUntil(self.clients.claim());
    });
})();
//# sourceMappingURL=serviceworker-v1.js.map
