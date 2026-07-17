export default function Comment({ comment }) {
	const date = new Date(comment.commentedAt).toLocaleString()

	return (
		<div>
			<div className="post-card-header">
				<h2>{comment.author}</h2>
				<span><small>{date}</small></span>
			</div>
			<div className="post-card-content">
				<p>{comment.content}</p>
			</div>
		</div>
	)
}
