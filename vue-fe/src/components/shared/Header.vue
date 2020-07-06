<template>
    <div class="vb-header">
        <div class="md-layout">
            <div class="md-layout-item">
                <p>Vue Blockchain header</p>
                <div>
                    <img alt="Vue logo" src="../../assets/logo.png">
                </div>
                <router-link to="/">Home</router-link>
                -
                <router-link to="/blockchain">Blockchain</router-link>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'Header'
    }
</script>

<style scoped>
    .vb-header {
        padding-bottom: 10px;
    }
</style>
