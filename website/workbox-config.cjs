// Workbox injectManifest 配置：在 `hugo` 构建出 public/ 之后运行 `npm run sw`。
// 它把 app shell 资源的「文件 + 内容版本号」注入 sw-src/serviceworker.js 里的 self.__WB_MANIFEST，
// 输出最终的 public/serviceworker-v1.min.js（与 html-head.html 注册的 URL 一致）。
module.exports = {
  globDirectory: 'public/',
  // 只预缓存 app shell；HTML 页面、5MB 的 search-data、SW 自身、workbox 库都不进预缓存。
  globPatterns: [
    'book.min.*.css',
    'prism.css',
    'prism.js',
    'flexsearch.min.js',
    'mermaid.min.js',
    '*.search.min.*.js',
    'manifest.json',
    'favicon.png',
    'logo.png',
    'apple-touch-icon-*.png',
    'fonts/*.woff2',
    'katex/katex.min.css',
    'katex/katex.min.js',
    'katex/auto-render.min.js',
    'katex/fonts/*.woff2',
  ],
  globIgnores: [
    'js/**',                       // 自托管的 workbox 运行时（importScripts 加载，不走预缓存）
    '**/sw.js',
    '**/sw.min.*.js',
    '**/serviceworker*.js',        // 不要把 SW 自己放进预缓存
    '**/*.search-data.min.*.js',   // 单个 ~2.7MB，太大，交给运行时缓存
  ],
  swSrc: 'sw-src/serviceworker.js',
  swDest: 'public/serviceworker-v1.min.js',
  // mermaid.min.js ~847KB，放宽上限以便纳入预缓存。
  maximumFileSizeToCacheInBytes: 3 * 1024 * 1024,
};
