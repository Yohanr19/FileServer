<script>

export default{
    data(){
        return {
            reports: null
        }
    },
    methods: {
        async fetchReports(){
            let res = await fetch('http://localhost:8000/api/reports')
            this.reports = await res.json()
        }
    },
    mounted(){
        this.fetchReports()
    }
}
</script>

<template>
    <h1>Reports</h1>
    <h1 v-if="!reports">Loading</h1>
    <table v-else class="report-table">
        <tr>
            <th>Filename</th>
            <th>Size(Bytes)</th>
            <th>Channel</th>
            <th>Subscribers</th>
            <th>Date</th>
            <th>Status</th>
        </tr>
        <tr v-for="report in reports" :id="report.id" class="report-row">
            <td>{{report.filename}}</td>
            <td>{{report.filesize}}</td>
            <td>{{report.channel}}</td>
            <td>{{report.subscriber_amount}}</td>
            <td>{{report.date}}</td>
            <td>{{report.status}}</td>
            
        </tr>
        </table>
</template>

<style>
table, th, td {
  border: 1px solid black;
}
</style>