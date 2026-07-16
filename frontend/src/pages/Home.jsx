import Greet from '../modules/Greet.jsx'
import Currently from '../modules/Currently.jsx'
import Social from '../modules/Social.jsx'
import LatestPost from '../modules/LatestPost.jsx'
import Gif from '../modules/Gif.jsx'

export default function Home() {
	return (
		<>
			<Greet />
			<div>
				<Currently />
				<Social />
			</div>
			<div>
				<LatestPost />
				<Gif url={"https://alvesafk.com/static/media/tux-linux.gif"}/>
			</div>
		</>
	);
}
