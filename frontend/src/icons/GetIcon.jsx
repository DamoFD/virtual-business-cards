import * as Icon from './Icon';

const getIcon = (name) => {
    const iconName = `${name.replace(/\s+/g, '')}Icon`;
    return Icon[iconName] || null;
}

export default getIcon
