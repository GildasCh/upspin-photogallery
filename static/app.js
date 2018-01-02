var app = new Vue({
    el: '#app',
    data: {
        imageCollection: files,
        theaterImage: '',
        showTheaterImage: false,
        number: 0,
        event: {}
    },
    created: function () {
        window.addEventListener('keydown', this.key);
    },
    methods: {
        show: function (imgNumber) {
            this.number = imgNumber;
            this.theaterImage = '/f/' + this.imageCollection[this.number];
            this.showTheaterImage = true;
        },
        key: function(keyEvent) {
            if (keyEvent.key == "ArrowLeft") {
                this.previous();
            } else if (keyEvent.key == "ArrowRight") {
                this.next();
            }
        },
        next: function () {
            this.number = (this.number + 1) % this.imageCollection.length;
            this.theaterImage = '/f/' + this.imageCollection[this.number];
        },
        previous: function () {
            this.number = (this.number - 1 + this.imageCollection.length) %
                this.imageCollection.length;
            this.theaterImage = '/f/' + this.imageCollection[this.number];
        },
        hide: function () {
            this.showTheaterImage = false;
        }
    }
})
