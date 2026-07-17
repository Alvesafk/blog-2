import { useState, useEffect } from 'react';
import { useLocation } from 'react-router';

export default function Post() {
	const [post, setPost] = useState(null);
	const [loading, setLoading] = useState(true);

	const location = useLocation()
	const { postId } = location.state || {};

	useEffect(() => {
		fetch("/api/post/" + postId)
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
				<article>
					<div dangerouslySetInnerHTML={{__html: post.content.content}}/>
				</article>
			) : (
				<p>No post found!</p>
			)}
		</>
	)
}
