<template>
    <div class="vb-blockchain">
        <div class="md-layout">
            <div class="md-layout-item md-small-hide"></div>
            <div class="md-layout-item md-small-100 vb-list">
                <md-card>
                    <md-card-header>
                        <div class="md-title">
                            {{title}}
                        </div>
                    </md-card-header>
                    <md-card-content v-if="blocks.length">
                        <div v-for="(block, id) in blocks" :key="id">
                            <Block :id="block.id"
                                   :nonce="block.nonce"
                                   :data="block.data"
                                   :prevHash="block.prevHash"
                                   :hash="block.hash">
                            </Block>
                        </div>
                    </md-card-content>
                    <md-card-content v-else>
                        <div>
                            <p>Empty blockchain</p>
                        </div>
                    </md-card-content>
                    <md-card-content>
                        <hr>
                        <div class="new-block-section">
                            <h2 class="new-block-title">Insert a new block</h2>
                            <md-field class="new-block-field">
                                <label>Data to insert</label>
                                <md-input v-model="newBlockData"></md-input>
                            </md-field>
                        </div>
                    </md-card-content>
                    <md-card-actions>
                        <md-button large class="md-dense md-primary md-theme-demo-light" @click="goHome">HOME
                        </md-button>
                        <md-button large class="md-dense md-raised md-primary md-theme-demo-light"
                                   @click="mineBlock(newBlockData)">Insert
                        </md-button>
                    </md-card-actions>
                </md-card>
            </div>
            <div class="md-layout-item md-small-hide"></div>
        </div>
    </div>
</template>

<script>
    import Block from './Block.vue'

    // Blockchain list
    export default {
        name: 'Blockchain',
        components: {
            Block
        },
        data() {
            return {
                title: 'Blockchain block list',
                newBlockData: '',
                blocks: []
            }
        },
        methods: {
            getBlockchain() {
                // Get blockchain list
                this.$utils.getAxios().get(this.$utils.getFullEndpoint('getBlockchain'))
                    .then(response => {
                        this.blocks = response.data.blockchain;
                    })
                    .catch(error => {
                        console.log(error);
                        alert('Error retrieving blockchain list :(');
                    })
            },
            mineBlock(data) {
                // Mine a new block
                this.$utils.getAxios().post(
                    this.$utils.getFullEndpoint('newBlock'), {
                        data
                    }, {
                        headers: this.$utils.getHttpHeaders(),
                    }
                ).then(response => {
                    console.log(response);
                    this.newBlockData = null;
                    this.getBlockchain();
                    alert('Block mined!');
                }).catch(error => {
                    console.log(error);
                    alert('Error inserting a new block :(');
                })
            },
            goHome() {
                this.$router.push('/');
            }
        },
        mounted() {
            this.getBlockchain()
        }
    }
</script>

<style scoped>
    .vb-blockchain {
        padding: 10px;
    }

    .vb-blockchain .vb-list {
        word-break: break-all;
    }
</style>
