import { defineStore } from 'pinia';

export const useMyUserStore = defineStore('myUserStore', {
  state: () => ({
    user: {
      id: '',
      username: '',
      email: '',
      role: ''
    },
    token: ''
  }),
  actions: {

  }
})
