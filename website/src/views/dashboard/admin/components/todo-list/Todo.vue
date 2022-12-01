<!--
 * @Description:  to do item
 * @Author: ZY
 * @Date: 2021-01-15 18:50:38
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2021-06-02 11:31:16
-->
<template>
  <li
    :class="{completed: todo.checked_at !== null, editing: editing}"
    class="todo"
  >
    <el-tooltip
      class="todo-tooltip"
      effect="dark"
      :content="todo.title"
      placement="top-start"
    >
      <div class="view">
        <input
          :checked="todo.checked_at !== null"
          class="toggle"
          type="checkbox"
          @change="toggleTodo(todo)"
        >
        <label
          style="padding-right: 40px;overflow:hidden;"
          class="text-ellipsis"
          @dblclick="editing = true"
          v-text="todo.title"
        />
        <button
          class="destroy"
          @click="deleteTodo( todo )"
        />
      </div>
    </el-tooltip>
    <input
      v-show="editing"
      v-focus="editing"
      :value="todo.title"
      class="edit"
      @keyup.enter="doneEdit"
      @keyup.esc="cancelEdit"
      @blur="doneEdit"
    >
  </li>
</template>

<script lang="ts">
import { defineComponent, PropType, ref } from 'vue'
import { TodoList } from '@/model/todoList'

export default defineComponent({
  name: 'Todo',
  directives: {
    focus: {
      // 指令的定义
      mounted(el) {
        el.focus()
      }
    }
  },
  props: {
    todo: {
      type: Object as PropType<TodoList>,
      default: () => {
        return {
          category: '',
          id: 0,
          title: ''
        } as TodoList
      }
    }
  },
  emits: ['toggle-todo', 'edit-todo', 'delete-todo'],
  setup(props, { emit }) {
    const editing = ref(false)
    const deleteTodo = (todo: TodoList) => {
      emit('delete-todo', todo)
    }
    const editTodo = ({ todo, value }: { todo: TodoList, value: string }) => {
      emit('edit-todo', { todo, value })
    }
    const toggleTodo = (todo: TodoList) => {
      emit('toggle-todo', todo)
    }

    const doneEdit = (e: KeyboardEvent) => {
      const value = (e.target as HTMLInputElement).value.trim()
      const todo = props.todo
      if (!value) {
        deleteTodo(todo)
      } else if (editing.value) {
        editTodo({
          todo,
          value
        })
        editing.value = false
      }
    }

    const cancelEdit = (e: KeyboardEvent) => {
      (e.target as HTMLInputElement).value = props.todo.title
      editing.value = false
    }

    return {
      editing,
      deleteTodo,
      editTodo,
      toggleTodo,
      doneEdit,
      cancelEdit
    }
  }
})
</script>
<style lang="scss" scoped>
.todo label {
  border-bottom: 1px solid #ddd;
}
</style>
