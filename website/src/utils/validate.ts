import { FormRulesMap } from 'element-plus/es/el-form/src/form.type'

export const isValidUsername = (str: string) => ['admin', 'editor'].indexOf(str.trim()) >= 0

export const isExternal = (path: string) => /^(https?:|mailto:|tel:)/.test(path)

export const isArray = (arg: any) => {
  if (typeof Array.isArray === 'undefined') {
    return Object.prototype.toString.call(arg) === '[object Array]'
  }
  return Array.isArray(arg)
}

export const isValidURL = (url: string) => {
  const reg = /^(https?|ftp):\/\/([a-zA-Z0-9.-]+(:[a-zA-Z0-9.&%$-]+)*@)*((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])){3}|([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{2}))(:[0-9]+)*(\/($|[a-zA-Z0-9.,?'\\+&%$#=~_-]+))*$/
  return reg.test(url)
}

export const getValidateFormDefaultRules = (formProps: string[]) : FormRulesMap => {
  const formRules : FormRulesMap = {}
  formProps.forEach((prop: string) => {
    formRules[prop] = []
  })
  for (let key in formRules) {
    if (Array.isArray(formRules[key]) && (formRules[key] as FormRulesMap[]).length === 0) {
      // default rule
      let isUser = false
      let isNumber = false
      let isArray = false
      let isTimestamp = false
      let isNotZero = false
      if (key.includes('.user')) {
        isUser = true
        key = key.replace('.user', '')
      }
      if (key.includes('.number')) {
        isNumber = true
        key = key.replace('.number', '')
      }
      if (key.includes('.nzero')) {
        isNotZero = true
        key = key.replace('.nzero', '')
      }
      if (key.includes('.array')) {
        isArray = true
        key = key.replace('.array', '')
      }
      if (key.includes('_at')) {
        isTimestamp = true
      }
      formRules[key] = [{
        required: true,
        trigger: 'blur',
        validator: function(rule, value, callback, source, options): boolean | Error | Error[] {
          const labelName = document.querySelector(`.el-form label[for="${rule.field}"]`)!.innerHTML.trim().replace(/(^\w、)|(:)/, '')
          const errors : Error[] = []
          if (value === null || value === undefined || value === '') {
            errors.push(new Error(`${labelName} 必填`))
            return errors
          }
          if (isArray) {
            try {
              const arr = JSON.parse(value)
              if (!Array.isArray(arr)) {
                errors.push(new Error(`${labelName} 只能为数组 格式:"[元素1,元素2,...]"`))
                return errors
              }
              if (isNumber && arr.every(item => !isNaN(value))) {
                errors.push(new Error(`${labelName} 只能为数组 格式:"[数字1,数字2,...]"`))
                return errors
              }
            } catch (e) {
              errors.push(new Error(`${labelName} 只能为数组 格式:"[元素1,元素2,...]"`))
              return errors
            }
            return errors
          }
          if (isNumber) {
            if (isNaN(value)) {
              errors.push(new Error(`${labelName} 只能为数字`))
              return errors
            }
          }
          if (isNotZero) {
            if (!isNaN(value) && Number(value) === 0) {
              errors.push(new Error(`${labelName} 不能为0`))
              return errors
            }
          }
          if (isUser && value === 0) {
            errors.push(new Error(`${labelName} 必填`))
            return errors
          }
          if (isTimestamp && value === 0) {
            errors.push(new Error(`${labelName} 必填`))
            return errors
          }
          return errors
        }
      }]
    }
  }
  return formRules
}
