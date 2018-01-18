Vue.config.devtools = true
var app = new Vue({
    el: '#app',
    data: {
        currentPage: ""
    },
    methods: {
        transEntry: function(event) {
            this.$http.get('/transEntry' ).then(function(response) {
                this.currentPage = response.data
                console.log("got transEntry")
            }).catch(function(error) {
                console.log(error)
            })
        },
        transList: function(event) {
            this.$http.get('/transList' ).then(function(response) {
                this.currentPage = response.data
                console.log("got transList")
            }).catch(function(error) {
                console.log(error)
            })
        } 
    }  
})