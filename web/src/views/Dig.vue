<script>
import DigService from "../services/DigService";
export default {
    name: "App",
    data() {
          return {
            query:{
              domain:'',
              record:''
            },
            records:'',
            showRecords:false
    
          }
    },
    watch:{
      query:{
        handler(newValue, oldValue) {
          console.log(1)
          this.showRecords=true
          this.sendQuery()

        },
        deep: true
      }
    },
    async created() {
          
    },
    methods: {
         async sendQuery() {
            const res = await DigService.digRecord(this.query);
            this.records =res.data.results
          }, 
          
    },
};
</script>
    
<template>
<div>
  <nav class="breadcrumb" aria-label="breadcrumbs">
    <ul>
      <li><a href="#">Code Challenge</a></li>
      <li class="is-active"><a href="#" aria-current="page">DIG</a></li>
    </ul>
  </nav>
  <div class="columns">
    <div class="column">
      <div class="box">
        <!--Domain-->
        <div class="field">
          <label class="label">Domain</label>
          <div class="control">
            <input class="input" type="text" placeholder="Text input" v-model="query.domain">
          </div>
        </div>
        <div class="field" v-show="query.domain!=''">
          <label class="label">Record type</label>
          <div class="control">
            <div class="select is-success">
              <select v-model="query.record">
                <option>A</option>
                <option>AAAA</option>
                <option>CNAME</option>
                <option>MX</option>
                <option>NS</option>
                <option>PTR</option>
                <option>SRV</option>
                <option>TXT</option>
                <option>SOA</option>
                <option>CAA</option>
                <option>TLSA</option>
                <option>TSIG</option>
              </select>
            </div>
          </div>
        </div>
      </div>
      <div class="box" v-show="showRecords">
        <article class="message" >
          <div class="message-header">
            <h1>Domain : {{query.domain}}</h1>
            
          </div>
          <div class="message-body" v-show="records">
            <h1>Record Type : {{query.record}}</h1>
            <div v-for="(record, index) in records" v-bind:key="index" v-show="records.length>0">
              <h1>{{record}}</h1>
            </div>
            <h1 v-show="records.length==0">Record not found !</h1>
          </div>
          
        </article>
      </div>
    </div>
      
  </div>
</div>
    
</template>
    