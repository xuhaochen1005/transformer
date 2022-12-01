import { checkPermission, checkRole } from '@/utils/permission';
export const role = {
    mounted(el, binding) {
        const { value } = binding;
        if (value && value instanceof Array && value.length > 0) {
            const hasPermission = checkRole(value);
            if (!hasPermission) {
                el.style.display = 'none';
            }
        }
        else {
            throw new Error('need roles! Like v-role="[\'admin\',\'editor\']"');
        }
    }
};
export const permission = {
    mounted(el, binding) {
        const { value } = binding;
        if (value && value instanceof Array && value.length > 1) {
            const hasPermission = checkPermission(value[0], value[1]);
            if (!hasPermission) {
                el.style.display = 'none';
            }
        }
        else {
            throw new Error('need roles! Like v-permission="[\'resource\', \'action\']"');
        }
    }
};
//# sourceMappingURL=index.js.map