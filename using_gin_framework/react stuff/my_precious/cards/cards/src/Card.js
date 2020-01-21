import React from 'react'


export default function Card() {
    return (
        <div className="col-lg-4">
            <div className="card mb-5 mb-lg-0">
            <div className="card-body">
                <h5 className="card-title text-muted text-uppercase text-center">Free</h5>
                <h6 className="card-price text-center">$0<span className="period">/month</span></h6>
                <hr />
                <ul>
                <li>Single User</li>
                <li>5GB Storage</li>
                <li>Unlimited Public Projects</li>
                <li>Community Access</li>
                <li className="text-muted">Unlimited Private Projects</li>
                <li className="text-muted">Dedicated Phone Support</li>
                <li className="text-muted">Free Subdomain</li>
                <li className="text-muted">Monthly Status Reports</li>
                </ul>
                <a href="https://github.com/EbookFoundation/free-programming-books" className="btn btn-block btn-primary text-uppercase">Button</a>
            </div>
            </div>
        </div>
    )
}
