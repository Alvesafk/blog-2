import { Link } from 'react-router';
import Tags from './Tags.jsx';

export default function PostCard({ post }) {
	const date = new Date(post.postedAt).toLocaleString()

	return (
		<div>
			<div className="post-card-header">
				<h2>{post.title}</h2>
				<span><small>{date}</small></span>
			</div>
			<div className="post-card-content">
				<p>{post.preview}</p>
				<Link to={"/posts/" + post.slugTitle}>Read it!</Link> <br />
				<Tags tags={post.tags} />
			</div>
		</div>
	)
}
