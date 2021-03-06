import React, { Component } from "react";
import Link from "gatsby-link";
import { ReactIcon as Icon } from "../svg/Icon";

class Logo extends Component {
  render() {
    const { handleClick } = this.props;
    return (
      <div>
        <Link
          to={"/"}
          style={{
            color: "black",
            textDecoration: "none",
            display: "block"
          }}
          onClick={handleClick}
        >
          <div
            style={{
              display: "flex",
              color: "white",
              fontSize: "1.5em",
              alignItems: "center"
            }}
          >
            <Icon />
            <div style={{ lineHeight: "60px" }}>React FAQ</div>
          </div>
        </Link>
      </div>
    );
  }
}

export default Logo;
