<template>
    <div class="vb-block">
        <hr>
        <div class="vb-block-item">
            <p><span class="header">Block id:</span> #{{id}}</p>
            <p><span class="header">Nonce:</span> {{nonce}}</p>
            <p><span class="header">Data:</span> {{data}}</p>
            <p><span class="header">prevHash:</span> {{prevHash}}</p>
            <p><span class="header">Hash:</span> {{hash}}</p>
        </div>
    </div>
</template>

<script>
    // Show block info
    export default {
        name: 'Block',
        props: {
            id: Number,
            nonce: Number,
            data: String,
            prevHash: String,
            hash: String,
        }
    }
</script>

<style scoped>
    .vb-block .vb-block-item span.header {
        font-weight: bold;
    }
</style>
