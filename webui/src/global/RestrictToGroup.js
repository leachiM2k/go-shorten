import React from 'react';

const isAllowed = neededGroups => {
    if (!Array.isArray(neededGroups) || neededGroups.length < 1) {
        return true;
    }

    const currentGroups = [];
    return neededGroups.some(neededGroup => currentGroups.includes(neededGroup));
};

const RestrictToGroup = (props) => {
    const { restrictToGroups, onFail } = props;

    if (!isAllowed(restrictToGroups)) {
        return onFail ? (<div id="ACCESS_FORBIDDEN">{onFail}</div>) : null;
    }
    return props.children;
};

export default RestrictToGroup;
export {isAllowed}
