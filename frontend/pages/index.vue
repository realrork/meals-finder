<template>
  <div class="flex flex-wrap gap-4 justify-center my-5">
    <div v-for="recipe in recipes" :key="recipe.id" class="card card-border bg-base-100 w-[18%] p-2">
      <NuxtImg :src="`/images/${recipe.id}.jpg`" alt="Recipe Image" class="w-auto rounded-lg"/>
      <p class="text-center h-16 font-semibold text-xl text-primary w-full mt-1 content-center">{{ recipe.name }}</p>
      <p class="text-center h-7 text-sm content-center">{{ recipe.description }}</p>
      <div class="card-actions justify-around mt-5">
        <button class="btn btn-secondary">Szczegóły</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const recipes = ref([
  {
    id: 1,
    name: 'Sezam chicken with with rice and vegetables',
    description: 'Sezam chicken with rice and vegetables',
  },
  {
    id: 2,
    name: 'Italian pancake of the forest smoothie',
    description: 'Italian pancake with cheese and vegetables AND with rice',
  },
  {
    id: 3,
    name: 'Fruits of the forest smoothie',
    description: 'Fruits of the forest smoothie',
  },
  {
    id: 4,
    name: 'Chicken curry',
    description: 'Chicken curry with rice',
  },
  {
    id: 5,
    name: 'Vegetable soup',
    description: 'Vegetable soup with bread',
  },
  {
    id: 6,
    name: 'Pasta with tomato sauce',
    description: 'Pasta with tomato sauce and cheese',
  },
  {
    id: 7,
    name: 'Vegetable stir-fry',
    description: 'Vegetable stir-fry with rice',
  },
  {
    id: 8,
    name: 'Chicken salad',
    description: 'Chicken salad with vegetables',
  },
  {
    id: 9,
    name: 'Fruit salad',
    description: 'Fruit salad with yogurt and honey',
  },
]);
</script>

<style></style>