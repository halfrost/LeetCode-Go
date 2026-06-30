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
precaching.precacheAndRoute([{"revision":"e6931bc096fc4f32f22e6faff4cd4569","url":"book.min.f8a130404c630c0548ad12d48a6b211d58536bddb42de301908d62338ec5bc78.css"},{"revision":"40847629b79d6bb9fab7af7d923d83bf","url":"prism.css"},{"revision":"5bc948ad60c30d12f8d6104c72e31aba","url":"prism.js"},{"revision":"ad47a5e1eedf6e613c404db19cd391c5","url":"flexsearch.min.js"},{"revision":"e7082840b9cf8a36e184c0f63bb474d6","url":"mermaid.min.js"},{"revision":"40e5e703e7b56fe02fb960c2dffbe0eb","url":"zh.search.min.057a4941d4e7b2be973e5deb120ed9fe048908cea71f272abe7c617f21c06caa.js"},{"revision":"1536bb0fe34850eb076a0b31a34c0ece","url":"en.search.min.cb4441a35f7edc8ad6327488282266c11589d9483728ec07a9653194baa463ff.js"},{"revision":"e10a3b32869a0b9e1e0fb7dd367b884d","url":"manifest.json"},{"revision":"393fb3b9b7ea8da7385140d19290990d","url":"favicon.png"},{"revision":"f0c6d1743ade4ec4fcf90c9545777392","url":"logo.png"},{"revision":"f7d6a4b65a8e7e67ea48cf9c683cb2a7","url":"apple-touch-icon-76x76.png"},{"revision":"c9ca6eb574e528519ad32128553eef69","url":"apple-touch-icon-60x60.png"},{"revision":"cb25a07256ee40552276dcfce0629edf","url":"apple-touch-icon-512x512.png"},{"revision":"da627e12f8328d6e5c506de7ca372325","url":"apple-touch-icon-192x192.png"},{"revision":"578c605f35b128a83d2856d4788af522","url":"apple-touch-icon-180x180.png"},{"revision":"3ce9ee0e58b2db681868dab6d0c46b36","url":"apple-touch-icon-152x152.png"},{"revision":"56ee88e2cff2308a391dacefd0127d19","url":"apple-touch-icon-120x120.png"},{"revision":"1d981480f686a38a34ff200bfa81ae79","url":"apple-touch-icon-1024x1024.png"},{"revision":"479970ffb74f2117317f9d24d9e317fe","url":"fonts/roboto-v19-latin-regular.woff2"},{"revision":"2735a3a69b509faf3577afd25bdf552e","url":"fonts/roboto-v19-latin-700.woff2"},{"revision":"14286f3ba79c6627433572dfa925202e","url":"fonts/roboto-v19-latin-300italic.woff2"},{"revision":"0c94e034ca06357576c2d03d623e1fcd","url":"fonts/roboto-mono-v6-latin-regular.woff2"},{"revision":"27f0bd8419d8acb9009ed4d9c35be068","url":"katex/katex.min.css"},{"revision":"c158c9e823b681cf535f46596b5e4eac","url":"katex/katex.min.js"},{"revision":"28cd0b98cd3f4fa37d52f3ffe47ad9d4","url":"katex/auto-render.min.js"},{"revision":"6cc31ea5c223c88705a13727a71417fa","url":"katex/fonts/KaTeX_Typewriter-Regular.woff2"},{"revision":"6a3255dfc1ba41c46e7e807f8ab16c49","url":"katex/fonts/KaTeX_Size4-Regular.woff2"},{"revision":"b311ca09df2c89a10fbb914b5a053805","url":"katex/fonts/KaTeX_Size3-Regular.woff2"},{"revision":"81d6b8d5ca77d63d5033d6991549a659","url":"katex/fonts/KaTeX_Size2-Regular.woff2"},{"revision":"048c39cba4dfb0460682a45e84548e4b","url":"katex/fonts/KaTeX_Size1-Regular.woff2"},{"revision":"755e2491f13b5269f0afd5a56f7aa692","url":"katex/fonts/KaTeX_Script-Regular.woff2"},{"revision":"d929cd671b19f0cfea55b6200fb47461","url":"katex/fonts/KaTeX_SansSerif-Regular.woff2"},{"revision":"fba01c9c6fb2866a0f95bcacb2c187a5","url":"katex/fonts/KaTeX_SansSerif-Italic.woff2"},{"revision":"6e0830bee40435e72165345e0682fbfc","url":"katex/fonts/KaTeX_SansSerif-Bold.woff2"},{"revision":"4ad08b826b8065e1eab85324d726538c","url":"katex/fonts/KaTeX_Math-Italic.woff2"},{"revision":"d747bd1e7a6a43864285edd73dcde253","url":"katex/fonts/KaTeX_Math-BoldItalic.woff2"},{"revision":"5c734d78610fa35282f3379f866707f2","url":"katex/fonts/KaTeX_Main-Regular.woff2"},{"revision":"e533d5a2506cf053cd671b335ec04dde","url":"katex/fonts/KaTeX_Main-Italic.woff2"},{"revision":"284a17fe5baf72ff8217d4c7e70c0f82","url":"katex/fonts/KaTeX_Main-BoldItalic.woff2"},{"revision":"8e1e01c4b1207c0a383d9a2b4f86e637","url":"katex/fonts/KaTeX_Main-Bold.woff2"},{"revision":"32a5339eb809f381a7357ba56f82aab3","url":"katex/fonts/KaTeX_Fraktur-Regular.woff2"},{"revision":"d5b59ec9764e10f4a82369ae29f3ac58","url":"katex/fonts/KaTeX_Fraktur-Bold.woff2"},{"revision":"7edb53b6693d75b8a2232481eea1a52c","url":"katex/fonts/KaTeX_Caligraphic-Regular.woff2"},{"revision":"4ec58befa687e9752c3c91cd9bcf1bcb","url":"katex/fonts/KaTeX_Caligraphic-Bold.woff2"},{"revision":"e78e28b4834954df047e4925e9dbf354","url":"katex/fonts/KaTeX_AMS-Regular.woff2"}] || []);
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
