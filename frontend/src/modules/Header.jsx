import { Link } from "react-router";

export default function Header() {
	return (
		<header>
			<h1>Alves's Blog</h1>

			<nav>
				<Link to="/">Home</Link>
				<Link to="/posts">Posts</Link>
				<Link to="/about">About</Link>
				<Link to="/resources">Resources</Link>
				<Link to="/guestbook">Guestbook</Link>
			</nav>
		</header>
	);
}
