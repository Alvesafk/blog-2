import { useState, useEffect } from 'react';
import { useParams } from 'react-router';
import ReactMarkdown from 'react-markdown';
import CommentSection from '../modules/CommentSection';

export default function Post() {
	const [post, setPost] = useState(null);
	const [loading, setLoading] = useState(true);

	const { name } = useParams()

	useEffect(() => {
		fetch("/api/post/" + name)
			.then(response => {
				if (response.ok) {
					return response.json();
				}
				throw response;
			})
			.then(p => {
				setPost(p);
				setLoading(false);
			})
			.catch(err => {
				console.error(`Error on fetching latest post: `, err);
				setLoading(false)
			});
	}, []);

	return (
		<>
			{loading ? (
				<p>Loading...</p>
			) : post ? (
				<>
					<article>
						<ReactMarkdown>{post.content.content}</ReactMarkdown>
					</article>
					<CommentSection id={post.content.id} />
				</>
			) : (
				<p>No post found!</p>
			)}
		</>
	)
}
