/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_API_BASE_URL: string
  readonly VITE_AWS_COGNITO_USER_POOL_ID: string
  readonly VITE_AWS_COGNITO_CLIENT_ID: string
  readonly VITE_AWS_S3_BUCKET: string
  readonly VITE_AWS_REGION: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
