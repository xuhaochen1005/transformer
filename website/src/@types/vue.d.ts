import { ElMessage } from 'element-plus'

declare module '@vue/runtime-core' {
  interface ComponentCustomProperties {
    $message: ElMessage
  }
}

declare module 'vue-router' {
  interface RouteMeta {
    title?: string
    roles?: string[]
    permission?: {
      resource: string
      action: string
    }
  }
}

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}
