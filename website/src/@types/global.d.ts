type func = () => void
interface Window {
  MathJax: {
    tex: {
      inlineMath: Array<Array<string>>
      displayMath: Array<Array<string>>
      processEnvironments: boolean
      processRefs: boolean
    }
    options: {
      skipHtmlTags: Array<string>,
      ignoreHtmlClass: string
    }
    startup: {
      pageReady: func
    }
    svg: {
      fontCache: string
    }
  }
}
