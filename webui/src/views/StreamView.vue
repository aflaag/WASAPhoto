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
            <p class="header">Your stream</p>

            <img class="user-icon" src="assets/user-small.svg">
        </div>

		<div class="horizontal-scroll-panel">
			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">walter4c</p>
				</div>

				<div class="post-photo-div">
					<img class="post-photo" src="assets/cupolone.jpg">
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

			<div class="post-card">
				<div class="post-card-header">
					<p class="post-card-username">enginemode1</p>
				</div>

				<div class="post-photo-div">
					<img class="post-photo" src="assets/salisburgo.jpg">
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
    .everything {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .header-div {
		display: flex;
        justify-content: space-between;

        margin-top: 1%;
        margin-left: 2%;
		margin-right: 2%;

        width: 100%;
    }

    .header {
        font-weight: 800;
        font-size: 500%;
        color: #485696;
    }

	.horizontal-scroll-panel {
		position: absolute;

		border: solid 3px red;

		top: 30%;
		left: 2%;

		width: 96%;
		height: 60%;

		overflow: auto;
		white-space: nowrap;
	}

	.post-card {
		display: inline-block;
		text-align: center;

		width: 25%;
		height: 100%;

		margin-right: 3%;

		border: solid 3px red;

        color: #485696;

		font-size: 300%;
	}

	.post-photo-div {
		width: 100%;
		height: 70%;

		margin-top: -2%;

		border: solid 3px blue;

	}

	.post-photo-div img {
		width: 100%;
		height: 100%;
		object-fit: contain;
		overflow: hidden;
	}

	.post-card-footer {
		display: flex;
		font-weight: 800;

		border: solid 1px green;
		height: 14%;
	}

	.post-photo-utils {
		border: solid 1px green;
		height: 90%;

		margin-left: 2%;
	}

	.post-card-footer p {
		margin-top: 1%;
		margin-right: 4%;
		margin-left: 2%;
	}

	.post-card-header {
		display: flex;
		margin-top: 2%;
	}

	.post-card-header p {
		margin-left: 6%;
	}
</style>
