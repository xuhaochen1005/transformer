import { ElLoading } from 'element-plus';
import { watch } from 'vue';
export function tLoading(loadingValue, target) {
    let loadingInstance = null;
    watch(loadingValue, function (value) {
        if (value) {
            loadingInstance = ElLoading.service({
                target: target
            });
        }
        else {
            loadingInstance.close();
        }
    });
}
//# sourceMappingURL=loading.js.map