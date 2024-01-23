<script>
	import CommentBox from "../components/CommentBox.vue";

    export default {
		components: { CommentBox },
        data: function() {
            return {
                errormsg: null,
                loading: false,

				token: localStorage.getItem("token"),
				uname: localStorage.getItem("uname"),

				comments: {},

				show_likes: false,
				likes: null,

				modal: null,

                show_follow: true,
                show_ban: true,

                photos: null,
                empty_photos: true,

                photo_count: 0,
                followers_count: 0,
                following_count: 0,

                show_followers: false,
                show_following: false,
            }
        },
        methods: {
			async getPhotoComments(photo) {
				try {
					let response = await this.$axios.get("/user/" + photo.user.username + "/photos/" + photo.id + "/comments", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

					this.comments = response.data;

					this.modal = new bootstrap.Modal(document.getElementById('logviewer'));
					this.modal.show();
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to fetch the photo's comments.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async updateLike(photo) {
				if (!photo.like_status) {
					try {
						let _ = await this.$axios.put("/user/" + photo.user.username + "/photos/" + photo.id + "/likes/" + this.uname, {}, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
						});

						photo.like_count += 1;
					} catch (e) {
						if (e.response && e.response.status === 500) {
							this.errormsg = "Something went wrong while trying to register the like.";
						} else if (e.response && e.response.status == 401) {
							this.errormgs = "Forbidden access"
						} else {
							this.errormsg = e.toString();
						}
					}
				} else {
					try {
						let _ = await this.$axios.delete("/user/" + photo.user.username + "/photos/" + photo.id + "/likes/" + this.uname, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
						});

						photo.like_count -= 1;
					} catch (e) {
						if (e.response && e.response.status === 500) {
							this.errormsg = "Something went wrong while trying to remove the like.";
						} else if (e.response && e.response.status == 401) {
							this.errormgs = "Forbidden access"
						} else {
							this.errormsg = e.toString();
						}
					}
				}

				photo.like_status = !photo.like_status;
			},
			async getPhotoLikes(photo) {
				try {
					let response = await this.$axios.get("/user/" + photo.user.username + "/photos/" + photo.id + "/likes", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

					this.likes = response.data;

					this.show_likes = true;

					if (this.likes.users.length > 0) {
						this.empty_likes = false;
					}
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve likes.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
			},
            async follow() {
                try {
					let _ = await this.$axios.put("/user/" + this.uname + "/follow/" + this.$route.params.uname, {}, {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.show_follow = !this.show_follow;

                    this.followers_count += 1;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to register the follow.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
            },
            async unfollow() {
                try {
					let _ = await this.$axios.delete("/user/" + this.uname + "/follow/" + this.$route.params.uname, {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.show_follow = !this.show_follow;

                    this.followers_count -= 1;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to remove the follow.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
            },
            async ban() {
                try {
					let _ = await this.$axios.put("/user/" + this.uname + "/ban/" + this.$route.params.uname, {}, {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.show_ban = !this.show_ban;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to register the ban.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
            },
            async unban() {
                try {
					let _ = await this.$axios.delete("/user/" + this.uname + "/ban/" + this.$route.params.uname, {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.show_ban = !this.show_ban;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to remove the ban.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
            },
			async home() {
                this.$router.push({path: "/user/" + this.uname + "/stream"});
			},
            async getProfileInfo() {
                try {
					let response = await this.$axios.get("/user/" + this.$route.params.uname, {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.show_follow = !response.data.follow_status;
                    this.show_ban = !response.data.ban_status;

                    this.photos = response.data.photos;

					if (this.photos.length > 0) {
						this.empty_photos = false;
					}

                    this.photo_count = response.data.photo_count;
                    this.followers_count = response.data.followers_count;
                    this.following_count = response.data.following_count;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve profile information.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
            },
            async getFollowers() {
                try {
					let response = await this.$axios.get("/user/" + this.$route.params.uname + "/followers", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.followers = response.data;

                    this.show_followers = true;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve the followers.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
            },
            async getFollowing() {
                try {
					let response = await this.$axios.get("/user/" + this.$route.params.uname + "/following", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.following = response.data;

                    this.show_following = true;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve the following.";
					} else if (e.response && e.response.status == 401) {
						this.errormgs = "Forbidden access"
					} else {
						this.errormsg = e.toString();
					}
				}
            }
        },
        mounted() {
            this.getProfileInfo()
        }
    }
</script>

<template>
    <div class="everything">
        <div class="header-div">
            <div class="username-div">
                <div style="display: flex; justify-content: center;">
                    <p class="header">{{this.$route.params.uname}}</p>
                </div>

                <div class="option-buttons-div">
                    <button v-if="!this.show_follow" @click="unfollow" class="button">
                        <img class="follow-icon" src="/assets/unfollow.svg">
                    </button>

                    <button v-if="this.show_follow" @click="follow" class="button">
                        <img class="follow-icon" src="/assets/follow.svg">
                    </button>

                    <button v-if="!this.show_ban" @click="unban" class="button">
                        <img class="ban-icon" src="/assets/unban.svg">
                    </button>

                    <button v-if="this.show_ban" @click="ban" class="button">
                        <img class="ban-icon" src="/assets/ban.svg">
                    </button>
                </div>
            </div>

            <div class="numbers-div">
                <div class="numbers-block-div">
                    <p class="numbers-block-numbers">{{this.photo_count}}</p>
                    <p style="margin-top: 3%">Photos</p>
                </div>

                <button @click="getFollowers" class="button">
                    <div class="numbers-block-div">
                        <p class="numbers-block-numbers">{{this.followers_count}}</p>
                        <p style="margin-top: 3%">Followers</p>
                    </div>
                </button>

                <button @click="getFollowing" class="button">
                    <div class="numbers-block-div">
                        <p class="numbers-block-numbers">{{this.following_count}}</p>
                        <p style="margin-top: 3%">Following</p>
                    </div>
                </button>
            </div>

            <button @click="home" class="button">
                <img class="user-icon" src="/assets/home.svg">
            </button>
        </div>

		<div v-if="!this.empty_photos" class="horizontal-scroll-panel">
			<div class="post-card" v-for="photo in this.photos" :key="photo.id">
				<div class="post-card-header" style="margin-top: 0px">
					<RouterLink :to="'/user/' + photo.user.username" class="nav-link" style="margin-left: 20px; margin-top: 6px; height: 80px;">
						<p class="post-card-username">{{photo.user.username}}</p>
					</RouterLink>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img class="post-photo-img" src="/assets/cupolone.jpg"> <!-- TODO: DA FARE -->
				</div>

				<div class="post-card-footer" style="margin-top: 7px">
					<button @click="updateLike(photo);" class="button" style="margin-bottom: 60px; margin-left: 20px">
						<div class="post-photo-utils" style="margin-right: 10px;">
							<img v-if="!photo.like_status" src="/assets/like-not-liked.svg"/>
							<img v-if="photo.like_status" src="/assets/like-liked.svg"/>
						</div>
					</button>

					<button @click="getPhotoLikes(photo)" class="button" style="margin: 5px 20px 0px 10px;">
						<p>{{photo.like_count}}</p>
					</button>

					<button @click="getPhotoComments(photo)" class="button" style="margin-bottom: 45px; margin-right: 8px">
						<div class="post-photo-utils">
							<img src="/assets/comment.svg"/>
						</div>
					</button>

					<p>{{photo.comment_count}}</p>

					<CommentBox id="logviewer" :comments="this.comments" :photo="photo" :modal="this.modal"></CommentBox>
				</div>
			</div>
		</div>

		<div v-if="this.empty_photos" class="horizontal-scroll-panel">
			<div style="display: flex; justify-content: center; margin-top: 13%">
				<p style="color: #485696; font-size: 300%">Find new users and follow your friends!</p>
			</div>
		</div>
    </div>

	<div v-if="this.show_likes" class="overlay" style="margin-top: -214px">
		<div class="comment-box">
			<button class="button" @click="this.show_likes = false;" style="display:flex">
				<img class="cross" src="/assets/cross.svg"/>
			</button>
						
			<div class="search-scroll-panel">
				<div v-for="like in this.likes.users" :key="like.id">
					<div class="comment">
						<div class="comment-header">
							<div class="comment-op">
								<RouterLink :to="'/user/' + like.username" class="nav-link">
									<p>{{like.username}}</p>
								</RouterLink>
							</div>
						</div>

						<div class="heightless-line"></div>
					</div>
				</div>
			</div>
		</div>
	</div>

    <div v-if="this.show_followers" class="overlay" style="margin-top: -214px">
		<div class="comment-box">
			<button class="button" @click="this.show_followers = false;" style="display:flex">
				<img class="cross" src="/assets/cross.svg"/>
			</button>
						
			<div class="search-scroll-panel">
				<div v-for="follower in this.followers.users" :key="follower.id">
					<div class="comment">
						<div class="comment-header">
							<div class="comment-op">
								<RouterLink @click="this.show_followers = false;" :to="'/user/' + follower.username" class="nav-link">
									<p>{{follower.username}}</p>
								</RouterLink>
							</div>
						</div>

						<div class="heightless-line"></div>
					</div>
				</div>
			</div>
		</div>
	</div>

    <div v-if="this.show_following" class="overlay" style="margin-top: -214px">
		<div class="comment-box">
			<button class="button" @click="this.show_following = false;" style="display:flex">
				<img class="cross" src="/assets/cross.svg"/>
			</button>
						
			<div class="search-scroll-panel">
				<div v-for="following in this.following.users" :key="following.id">
					<div class="comment">
						<div class="comment-header">
							<div class="comment-op">
								<RouterLink @click="this.show_following = false;" :to="'/user/' + following.username" class="nav-link">
									<p>{{following.username}}</p>
								</RouterLink>
							</div>
						</div>

						<div class="heightless-line"></div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style>
</style>