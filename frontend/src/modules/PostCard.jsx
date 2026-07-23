import { Link } from 'react-router';
import Tags from './Tags.jsx';

export default function PostCard({ post }) {
	const date = new Date(post.postedAt).toLocaleString()

	return (
		<section className='container mt-2'>
			<div className='box'>
				<header className='has-b-border'>
					<h2 className='is-size-4'>{post.title}</h2>
					<span><small>{date}</small></span>
				</header>
				<div className="post-card-content">
					<p>{post.preview}</p>
					<Link to={"/posts/" + post.slugTitle}>Read it!</Link> <br />
					<Tags tags={post.tags} />
				</div>
			</div>
		</section>
	)
}
