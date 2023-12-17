<script>
    export default {
        data: function() {
            return {
                errormsg: null,
                loading: false,
                some_data: null,
            }
        },
        methods: {
            async refresh() {
                this.loading = true;
                this.errormsg = null;
                try {
                    let response = await this.$axios.get("/session");
                    this.some_data = response.data;
                } catch (e) {
                    this.errormsg = e.toString();
                }
                this.loading = false;
            },
        },
        mounted() {
            this.refresh()
        }
}
</script>

<template>
    <div class="everything">
        <header>
            <p class="header">WASAPhoto</p>
        </header>

        <div class="center-div">
            <div class="user-icon-div">
                <img class="user-icon" src="assets/user.svg">
            </div>

            <div class="bar-section-div">
                <input class="bar" placeholder="Enter your username!">

                <button class="button">boh</button>
            </div>
        </div>
    </div>
</template>

<style>
    .everything {
        display: flex;
        justify-content: center;
        align-items: center;

        width: 1920px;
        height: 1080px;

        margin: 0px;
    }

    .header {
        font-weight: 800;
        font-size: 250%;

        color: #485696;
    }
</style>