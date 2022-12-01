import { ILoadingInstance } from 'element-plus/packages/loading/src/loading.type'
import { ElLoading } from 'element-plus'
import { Ref, watch } from 'vue'

export function tLoading(loadingValue : Ref<boolean>, target: string) {
  let loadingInstance : ILoadingInstance | null = null
  watch(loadingValue, function(value) {
    if (value) {
      loadingInstance = ElLoading.service({
        target: target
      })
    } else {
      loadingInstance!.close()
    }
  })
}
