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
                    let response = await this.$axios.get("/stream"); // TODO: DA CAMBIARE
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
        <div class="header-div">
            <div class="username-div">
                <div style="display: flex; justify-content: center;">
                    <p class="header">enginemode1</p>
                </div>

                <div class="option-buttons-div">
                    <img class="follow-icon" src="assets/follow.svg">

                    <img class="ban-icon" src="assets/ban.svg">
                </div>
            </div>

            <div class="numbers-div">
                <div class="numbers-block-div">
                    <p class="numbers-block-numbers">52</p>
                    <p>Photos</p>
                </div>

                <div class="numbers-block-div">
                    <p class="numbers-block-numbers">0</p>
                    <p style="margin-top: 3%">Followers</p>
                </div>

                <div class="numbers-block-div">
                    <p class="numbers-block-numbers">1504</p>
                    <p style="margin-top: 3%">Following</p>
                </div>
            </div>

            <img class="user-icon" src="assets/home.svg">
        </div>

		<div class="horizontal-scroll-panel">
			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">walter4c</p>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="assets/cupolone.jpg">
				</div>

				<div class="post-card-footer">
					<div class="post-photo-utils">
						<img src="assets/like-liked.svg">
					</div>

					<p>2.6k</p>

					<div class="post-photo-utils">
						<img src="assets/comment.svg">
					</div>

					<p>1.5k</p>
				</div>
			</div>

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">walter4c</p>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="assets/salisburgo.jpg">
				</div>

				<div class="post-card-footer">
					<div class="post-photo-utils">
						<img  src="assets/like-liked.svg">
					</div>

					<p>2.6k</p>

					<div class="post-photo-utils">
						<img src="assets/comment.svg">
					</div>

					<p>1.5k</p>
				</div>
			</div>
		</div>
    </div>
</template>

<style>
</style>