var app = new Vue({
    el: '#app',
    data: {
        theaterImage: '',
        showTheaterImage: false
    },
    methods: {
        show: function (imgPath) {
            this.theaterImage = imgPath;
            this.showTheaterImage = true;
        },
        hide: function () {
            this.showTheaterImage = false;
        }
    }
})
