/*
 * @Description:全局存储类封装（AES加解密）
 * @Autor: ZY
 * @Date: 2020-11-04 11:51:01
 * @LastEditors: ZY
 * @LastEditTime: 2020-12-28 16:57:57
 */
import Cookies from 'js-cookie';
export var StorageType;
(function (StorageType) {
    StorageType[StorageType["local"] = 0] = "local";
    StorageType[StorageType["session"] = 1] = "session";
    StorageType[StorageType["cookie"] = 2] = "cookie";
})(StorageType || (StorageType = {}));
class VStorage {
    static shared() {
        if (!this.instance) {
            this.instance = new VStorage();
        }
        return this.instance;
    }
    /**
     * @description: 本地保存数据AES加密处理
     * @param {StorageType} type localStorage 和 sessionStorage 选择
     * @param {string} key
     * @param {T} value
     * @return {*}
     */
    rcSetItem(type = StorageType.local, key, value) {
        const valueJson = JSON.stringify(value);
        const valueAes = valueJson;
        if (type === StorageType.local) {
            localStorage.setItem(key, valueAes);
        }
        else if (type === StorageType.session) {
            sessionStorage.setItem(key, valueAes);
        }
        else {
            Cookies.set(key, valueAes);
        }
    }
    rcGetItem(type = StorageType.local, key) {
        if (type === StorageType.local) {
            return localStorage.getItem(key);
        }
        else if (type === StorageType.session) {
            return sessionStorage.getItem(key);
        }
        else {
            return Cookies.get(key);
        }
    }
    rcRemoveItem(type = StorageType.local, key) {
        if (type === StorageType.local) {
            localStorage.removeItem(key);
        }
        else if (type === StorageType.session) {
            sessionStorage.removeItem(key);
        }
        else {
            Cookies.remove(key);
        }
    }
}
export default VStorage.shared();
//# sourceMappingURL=storage.js.map