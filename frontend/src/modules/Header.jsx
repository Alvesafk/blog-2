import { Link } from "react-router";

export default function Header() {
	return (
		<header>
			<h1>Alves's Blog</h1>

			<nav>
				<Link to="/" prefetch="intent">Home</Link>
				<Link to="/posts" prefetch="intent">Posts</Link>
				<Link to="/resources">Resources</Link>
				<Link to="/guestbook" prefetch="intent">Guestbook</Link>
			</nav>
		</header>
	);
}
