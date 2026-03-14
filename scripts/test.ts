/**
 * 小安计划 API 测试脚本 (TypeScript)
 * 使用方式: 设置环境变量后运行 npx tsx scripts/test.ts
 *   BASE_URL=https://your-api.com EMAIL=xxx@example.com PASSWORD=xxx npx tsx scripts/test.ts
 */

const env = (key: string, defaultVal: string): string =>
  process.env[key] ?? defaultVal;

const baseURL = env('BASE_URL', 'http://localhost:3000');
const email = env('EMAIL', '');
const password = env('PASSWORD', '');

interface CommonResponse<T> {
  code: number;
  message: string;
  data: T;
}

async function request<T>(
  method: string,
  path: string,
  body?: object,
  token?: string
): Promise<CommonResponse<T>> {
  const url = path.startsWith('http') ? path : `${baseURL}${path}`;
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
  };
  if (token) headers['Authorization'] = `Bearer ${token}`;

  const res = await fetch(url, {
    method,
    headers,
    body: body ? JSON.stringify(body) : undefined,
  });
  const data = await res.json();
  if (!res.ok) throw new Error(`HTTP ${res.status}: ${JSON.stringify(data)}`);
  return data as CommonResponse<T>;
}

async function main() {
  if (!email || !password) {
    console.error('请设置环境变量: EMAIL, PASSWORD (可选 BASE_URL)');
    process.exit(1);
  }

  console.log(`[1/4] 连接: ${baseURL}`);

  // 1. 登录
  console.log('[2/4] 登录...');
  const loginResp = await request<{ token: string }>('POST', '/api/auth/login', {
    email,
    password,
    type: 'email',
  });
  if (loginResp.code !== 0) {
    console.error('登录失败:', loginResp.code, loginResp.message);
    process.exit(1);
  }
  const token = loginResp.data.token;
  console.log('登录成功, token 已设置');

  // 2. 获取用户信息
  console.log('[3/4] 获取用户信息...');
  const userResp = await request<{
    user_id: number;
    email: string;
    nickname: string;
    avatar: string;
    role: string;
  }>('GET', '/api/auth/get-user-info?user_id=0', undefined, token);
  if (userResp.code !== 0) {
    console.error('获取用户信息失败:', userResp.code, userResp.message);
    process.exit(1);
  }
  console.log('用户信息:', JSON.stringify(userResp.data, null, 2));

  // 3. 获取最新文章列表
  console.log('[4/4] 获取最新文章...');
  const articlesResp = await request<{ list: unknown[]; total: number }>(
    'GET',
    '/api/content/get-new-articles?page_size=10&cursor=0',
    undefined,
    token
  );
  if (articlesResp.code !== 0) {
    console.error('获取文章列表失败:', articlesResp.code, articlesResp.message);
    process.exit(1);
  }
  console.log('文章数量:', articlesResp.data.total);

  console.log('\n全部测试通过.');
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
