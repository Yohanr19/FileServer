<script>

export default{
    data(){
        return {
            status: null
        }
    },
    methods: {
        async fetchStatus(){
            this.status = null
            let res = await fetch('http://localhost:8000/api/status')
            this.status = await res.json()
        }
    },
    mounted(){
        this.fetchStatus()
    }
}
</script>

<template>
    <h1>Active Channels</h1>
    <h1 v-if="!status">Loading...</h1>
    <div v-else class="ch-wrapper">
        <div v-for="item in status" :key="item.channel" class="item">
            <h2>Channel: {{item.channel}}</h2>
            <h2>Listeners: {{item.amount}}</h2>
            <h2>Addresses: </h2>
        <div v-for="addr in item.subscriber_addrs" :key="addr">
            {{addr}}
        </div>
        </div>
    </div>
   
</template>

<style>
 .ch-wrapper{
    display:flex
 }
 .item{
    padding: 5px;
    border: 10px;
 }
</style>