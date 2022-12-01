/*
 * @Description: 根绝大小变化重新布局
 * @Author: ZY
 * @Date: 2020-12-17 15:37:56
 * @LastEditors: ZY
 * @LastEditTime: 2021-01-28 16:29:49
 */
// refer to Bootstrap's responsive design
import { useStore } from '@/store';
import { AppMutation } from '@/store/app/mutation';
import { DeviceType } from '@/store/state';
import { computed, watch } from 'vue';
import { useRoute } from 'vue-router';
const store = useStore();
const WIDTH = 992; // refer to Bootstrap's responsive design
export default function () {
    const device = computed(() => {
        return store.state.app.device;
    });
    const sidebar = computed(() => {
        return store.state.app.sidebar;
    });
    const currentRoute = useRoute();
    const watchRouter = watch(() => currentRoute.name, () => {
        if (store.state.app.device === DeviceType.Mobile && store.state.app.sidebar.opened) {
            store.commit(AppMutation.CLOSE_SIDEBAR, false);
        }
    });
    const isMobile = () => {
        const rect = document.body.getBoundingClientRect();
        return rect.width - 1 < WIDTH;
    };
    const resizeMounted = () => {
        if (isMobile()) {
            store.commit(AppMutation.TOGGLE_DEVICE, DeviceType.Mobile);
            store.commit(AppMutation.CLOSE_SIDEBAR, true);
        }
    };
    const resizeHandler = () => {
        if (!document.hidden) {
            store.commit(AppMutation.TOGGLE_DEVICE, isMobile() ? DeviceType.Mobile : DeviceType.Desktop);
            if (isMobile()) {
                store.commit(AppMutation.CLOSE_SIDEBAR, true);
            }
        }
    };
    const addEventListenerOnResize = () => {
        window.addEventListener('resize', resizeHandler);
    };
    const removeEventListenerResize = () => {
        window.removeEventListener('resize', resizeHandler);
    };
    return {
        device,
        sidebar,
        resizeMounted,
        addEventListenerOnResize,
        removeEventListenerResize,
        watchRouter
    };
}
//# sourceMappingURL=resize.js.map