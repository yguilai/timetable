<template>
  <div id="app">
    <transition name="app-fade">
      <router-view />
    </transition>
  </div>
</template>

<style lang="scss">
#app {
  width: 100vw;
  height: 100vh;
}

.app-fade-enter-active, .fade-leave-active {
  transition: opacity .3s
}

.app-fade-leave-to {
 display: none
}

.app-fade-enter, .app-fade-leave-to /* .app-fade-leave-active below version 2.1.8 */ {
  opacity: 0
}
</style>
