import ElementPlus, { ElMessage } from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'
import locale from 'element-plus/lib/locale/lang/zh-cn'

export default function setupElementPlus(app: any): void {
  app.use(ElementPlus, { locale: locale })
  app.config.globalProperties.$message = ElMessage
}
