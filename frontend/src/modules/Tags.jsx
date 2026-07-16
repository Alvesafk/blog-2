export default function Tags({ tags }) {
	let result = "";
	tags.forEach(function(element, index, array) {
		if (index == array.length - 1) {
			result += element;
		} else {
			result += element + ", ";
		}
	});

	return (
		<span>
			<small>
				{result}
			</small>
		</span>
	);
}
