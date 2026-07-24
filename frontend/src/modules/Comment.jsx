export default function Comment({ comment }) {
	const date = new Date(comment.commentedAt).toLocaleString()

	return (
		<section className="container is-max-desktop my-2">
			<div className="box has-border">
				<div className="post-card-header has-b-border">
					<h2 className="is-size-5">{comment.author}</h2>
					<span><small className="is-size-7">{date}</small></span>
				</div>
				<div className="post-card-content my-2">
					<p>{comment.content}</p>
				</div>
			</div>
		</section>
	)
}
