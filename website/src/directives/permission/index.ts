/*
 * @Description: 权限指令
 * @Author: ZY
 * @Date: 2020-12-28 10:39:21
 * @LastEditors: ZY
 * @LastEditTime: 2020-12-28 13:46:23
 */
import { Directive } from 'vue'
import { checkPermission, checkRole } from '@/utils/permission'

export const role: Directive = {
  mounted(el, binding) {
    const { value } = binding
    if (value && value instanceof Array && value.length > 0) {
      const hasPermission = checkRole(value)
      if (!hasPermission) {
        el.style.display = 'none'
      }
    } else {
      throw new Error('need roles! Like v-role="[\'admin\',\'editor\']"')
    }
  }
}

export const permission: Directive = {
  mounted(el, binding) {
    const { value } = binding
    if (value && value instanceof Array && value.length > 1) {
      const hasPermission = checkPermission(value[0], value[1])
      if (!hasPermission) {
        el.style.display = 'none'
      }
    } else {
      throw new Error('need roles! Like v-permission="[\'resource\', \'action\']"')
    }
  }
}
