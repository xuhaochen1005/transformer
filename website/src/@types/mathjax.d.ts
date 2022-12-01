declare module 'mathjax-vue3' {
  import type { DefineComponent } from 'vue'
  export const MathJax: DefineComponent<{}, {}, any>
  export function initMathJax(a: any, b: any): void
}
