import { useState, useEffect } from 'react'
import TextAreaForm from './TextAreaForm.jsx'
import Comment from './Comment.jsx'

export default function CommentSection({ id }) {
	const [comments, setComments] = useState(null);
	const [loading, setLoading] = useState(true);


	function fetchComments() {
		setLoading(true)

		fetch("/api/post/comments/" + id)
			.then(response => {
				if (response.ok) {
					return response.json();
				}
				throw response;
			})
			.then(c => {
				setComments(c);
				setLoading(false);
			})
			.catch(err => {
				console.error(`Error on fetching latest post: `, err);
				setLoading(false)
			});
	}
	useEffect(() => {
		fetchComments()
	}, [id]);

	return (
		<>
			<hr />
			<div id='comment-section'>
				<TextAreaForm
					url={"/api/post/" + id}
					onPost={fetchComments}
					inputPlaceholder={"Write a comment!"}
					submitPlaceholder={"Post your comment!"}
				/>
			</div>

			{loading ? (
				<p>Loading...</p>
			) : comments.content !== null ? (
				<>
					{comments.content.map(c =>
						<Comment key={c.id} comment={c} />
					)}
				</>
			) : (
				<p>No comments <small>yet...</small></p>
			)}
		</>
	)
}
