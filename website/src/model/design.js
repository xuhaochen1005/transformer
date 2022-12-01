// 任务状态
export var DesignTaskStatus;
(function (DesignTaskStatus) {
    DesignTaskStatus[DesignTaskStatus["DesignTaskCanceled"] = 1] = "DesignTaskCanceled";
    DesignTaskStatus[DesignTaskStatus["DesignTaskCreated"] = 2] = "DesignTaskCreated";
    DesignTaskStatus[DesignTaskStatus["DesignTaskStarted"] = 3] = "DesignTaskStarted";
    DesignTaskStatus[DesignTaskStatus["DesignTaskFinished"] = 4] = "DesignTaskFinished";
    DesignTaskStatus[DesignTaskStatus["DesignTaskReviewUnaccepted"] = 5] = "DesignTaskReviewUnaccepted";
    DesignTaskStatus[DesignTaskStatus["DesignTaskReviewed"] = 6] = "DesignTaskReviewed";
    DesignTaskStatus[DesignTaskStatus["DesignTaskApproveUnaccepted"] = 7] = "DesignTaskApproveUnaccepted";
    DesignTaskStatus[DesignTaskStatus["DesignTaskApproved"] = 8] = "DesignTaskApproved";
})(DesignTaskStatus || (DesignTaskStatus = {}));
// 接缝形式
export var SeamType;
(function (SeamType) {
    SeamType[SeamType["FullOblique"] = 1] = "FullOblique";
})(SeamType || (SeamType = {}));
//# sourceMappingURL=design.js.map