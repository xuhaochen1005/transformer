import request from '@/utils/request';
export function getDesignTaskList(params) {
    return request({
        url: '/design/tasks',
        method: 'GET',
        params: params
    });
}
export function createDesignTask(params) {
    return request({
        url: '/design/task',
        method: 'POST',
        data: params
    });
}
export function updateDesignTask(params) {
    return request({
        url: '/design/task/' + params.id,
        method: 'PATCH',
        data: params
    });
}
export function deleteDesignTask(id) {
    return request({
        url: '/design/task/' + id,
        method: 'DELETE'
    });
}
export function getRecommandDesignTaskList() {
    return request({
        url: '/design/recommand',
        method: 'GET'
    });
}
//# sourceMappingURL=design.js.map