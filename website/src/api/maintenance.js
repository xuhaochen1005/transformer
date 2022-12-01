import request from '@/utils/request';
export function GetSysHost(param) {
    return request({
        url: '/sys/host',
        method: 'GET',
        data: param
    });
}
export function GetSysCPU(param) {
    return request({
        url: '/sys/cpu',
        method: 'GET',
        data: param
    });
}
export function GetSysNet(param) {
    return request({
        url: '/sys/net',
        method: 'GET',
        data: param
    });
}
export function GetSysProcess(param) {
    return request({
        url: '/sys/process',
        method: 'GET',
        data: param
    });
}
export function GetSysMemory(param) {
    return request({
        url: '/sys/mem',
        method: 'GET',
        data: param
    });
}
export function GetSysBasic(param) {
    return request({
        url: '/sys/basic',
        method: 'GET',
        data: param
    });
}
export function GetSysDisk(param) {
    return request({
        url: '/sys/disk',
        method: 'GET',
        data: param
    });
}
//# sourceMappingURL=maintenance.js.map