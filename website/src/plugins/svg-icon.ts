import SvgIcon from '@/components/svg-icon/index.vue'

export default function setupSvgIcon(app: any) {
  const ctx = require.context('@/assets/svg', false, /\.svg$/)
  ctx.keys().forEach(path => {
    const tmp = path.match(/\.\/([A-Za-z0-9\-_]+)\.svg$/)
    if (!tmp) {
      return
    }
    const name = tmp[1]
    require(`@/assets/svg/${name}.svg`)
  })
  app.component(SvgIcon.name, SvgIcon)
}
