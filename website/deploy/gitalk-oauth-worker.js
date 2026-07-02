// Gitalk OAuth 代理 —— Cloudflare Worker 版
// 用途：替代 nginx 直连 GitHub 的方案。当「服务器到 github.com 出网不稳定」(随机 502/504/404)时，
//       把 token 交换放到 Cloudflare 边缘（到 github 稳定），彻底绕开你服务器的网络问题。
//
// 部署（二选一）：
//   A. 仪表盘最简单：
//      1. Cloudflare 控制台 → Workers & Pages → Create → Create Worker
//      2. 把本文件内容整段贴进去 → Deploy
//      3. 记下分配的地址，形如 https://gitalk-oauth.<你的子域>.workers.dev
//   B. 命令行：npm i -g wrangler && wrangler deploy
//
// 部署后改一行 config.toml：
//   [params.gitalk]
//     proxy = "https://gitalk-oauth.<你的子域>.workers.dev"
//   然后重新构建发布（npm run build）。
//
// 注意：Worker 与站点不同源，必须返回 CORS 头（下面已处理）；这点和原来的同源 nginx 代理不同。

const GITHUB_TOKEN_URL = 'https://github.com/login/oauth/access_token';

function corsHeaders() {
  return {
    'Access-Control-Allow-Origin': '*',
    'Access-Control-Allow-Methods': 'POST, OPTIONS',
    'Access-Control-Allow-Headers': 'Content-Type, Accept',
  };
}

export default {
  async fetch(request) {
    // 预检请求
    if (request.method === 'OPTIONS') {
      return new Response(null, { headers: corsHeaders() });
    }
    if (request.method !== 'POST') {
      return new Response(JSON.stringify({ error: 'method_not_allowed' }), {
        status: 405,
        headers: { 'Content-Type': 'application/json', ...corsHeaders() },
      });
    }

    const body = await request.text();
    const ghResp = await fetch(GITHUB_TOKEN_URL, {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': request.headers.get('Content-Type') || 'application/json',
        'User-Agent': 'gitalk-oauth-worker',
      },
      body,
    });

    const respBody = await ghResp.text();
    return new Response(respBody, {
      status: ghResp.status,
      headers: { 'Content-Type': 'application/json', ...corsHeaders() },
    });
  },
};
