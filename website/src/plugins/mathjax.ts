import MathJax from 'mathjax-vue3'

export default function setupMathJax(app: any): void {
  window.MathJax = {
    tex: {
      inlineMath: [['$', '$']],
      displayMath: [['$$', '$$']],
      processEnvironments: true,
      processRefs: true
    },
    options: {
      skipHtmlTags: ['noscript', 'style', 'textarea', 'pre', 'code'],
      ignoreHtmlClass: 'tex2jax_ignore'
    },
    startup: {
      pageReady: () => {
        console.log('Math Jax 初始化完毕')
      }
    },
    svg: {
      fontCache: 'global'
    }
  }
  app.use(MathJax)
}
