
const Message = ({messages}) => {
    const list = messages.map(message => {
        return (
            <li>
                {message.Title}
            </li>
        )
    })
    return <ul>{list}</ul>
}

export default Message