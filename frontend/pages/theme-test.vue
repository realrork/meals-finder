<template>
  <div>
    <!-- <div v-for="item in themes" :key="item">
      <input type="radio" v-model="selectedTheme" :value="item" v-bind:key="item" class="mr-2 " />{{ item }}
    </div> -->
    <select class="select" v-model="selectedTheme">
      <option  v-for="item in themes" :key="item" :value="item" class="mr-2 ">{{ item }}</option>
    </select>
    
  </div>
</template>

<script lang="ts" setup>
const themes = ref([
  'dracula',
  'dark',
  'light',
  'cupcake',
  'bumblebee',
  'emerald',
  'corporate',
  'synthwave',
  'retro',
  'cyberpunk',
  'valentine',
  'halloween',
  'garden',
  'forest',
  'aqua',
  'lofi',
  'pastel',
  'fantasy',
  'wireframe',
  'black',
  'luxury',
  'cmyk',
  'autumn',
  'business',
  'acid',
  'lemonade',
  'night',
  'coffee',
  'winter',
  'dim',
  'nord',
  'sunset',
  'caremellatte',
  'abbyss',
  'silk'
]);
const selectedTheme = ref('dracula');
watch(selectedTheme, () => {
  useHead({
    htmlAttrs: {
      'data-theme': selectedTheme.value
    },
  })
});

</script>

<style></style>