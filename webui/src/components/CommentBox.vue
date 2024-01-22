<script>
export default {
	props: ["log","token"],
	data() {
		return {
		}
	},

	methods: {
		async deleteComment(username, photoid, commentid ) {
			try {
				let response = await this.$axios.delete("/users/" + username + "/photo/" + photoid + "/comment/" + commentid, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("token")
					}
				})
				location.reload();
			} catch(e) {
				if (e.response && e.response.status === 400) {
                    this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
					this.detailedmsg = null;
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
					this.detailedmsg = e.toString();
				} else {
					this.errormsg = e.toString();
					this.detailedmsg = null;
				}
			}
		},
	},
	mounted() {
    }
}
</script>

<template>
	<div class="modal modal-xl" tabindex="-1">
		<div class="modal-dialog comment-box" style="margin-top: 170px">
			<div class="modal-content">
				<div class="modal-header">
                    <button class="button" style="display:flex;" data-bs-dismiss="modal">
                        <img class="cross" src="/assets/cross.svg"/>
                    </button>
				</div>
				<div class="modal-body" >
					<div class="comment-scroll-panel" style="height: 500px; position: absolute; margin-top: -50px; margin-left: -15px">
						<div class="comment" v-for="comment in log" :key="comment.id">
                            <div class="comment-header">
                                <div class="comment-op">
                                    <p>enginemode1</p>
                                </div>

                                <img class="delete-comment" src="/assets/delete.svg"/>
                            </div>

                            <div class="comment-text">
                                <p>lesgoooooooooooooooooooooooooooooooooooooooo ooooooooooooooooooooooooooooooooooooooooooooooooooooooo ooooooooooooooooooooooooooooooooooo</p>
                            </div>

                            <div class="heightless-line"></div>
						</div>
					</div>
				</div>

                <div class="comment-input-box">
                    <input class="comment-bar" placeholder="Leave a comment!">
                </div>
			</div>
		</div>
	</div>
</template>

<style>
    .modal-content {
		position: relative;

		width: 100%;
		height: 100%;

		border-radius: 30px;

		background-color: #e7e7e7;

		font-size: 90%;

		z-index: 1;
    }
</style>