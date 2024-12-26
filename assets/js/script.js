document.addEventListener("DOMContentLoaded", function () {
    const popupForm = document.getElementById("adoptFormPopup");
    const petNameSpan = document.getElementById("petName");
    const path = window.location.pathname;
    const petName = path.split("/").pop();

    if (petNameSpan) {
        petNameSpan.textContent = petName;
    }

    const adoptBtn = document.getElementById("adoptBtn");
    if (adoptBtn && popupForm) {
        adoptBtn.addEventListener("click", function () {
            popupForm.style.display = "flex";
        });
    }

    const closeBtn = document.querySelector(".close-btn");
    if (closeBtn) {
        closeBtn.addEventListener("click", function () {
            popupForm.style.display = "none";
        });
    }

    if (popupForm) {
        popupForm.addEventListener("click", function (e) {
            if (e.target === popupForm) {
                popupForm.style.display = "none";
            }
        });
    }

    // Tombol Approve dan Reject
    const approveButtons = document.querySelectorAll(".approve-btn");
    const rejectButtons = document.querySelectorAll(".reject-btn");

    approveButtons.forEach(btn => {
        btn.addEventListener("click", function () {
            const id = this.getAttribute("data-id");
            fetch(`/dashboard/adoptions/${id}/approve`, { method: "POST" })
                .then(response => {
                    if (response.ok) {
                        return response.json();
                    } else {
                        throw new Error("Failed to approve adoption");
                    }
                })
                .then(data => {
                    alert("Adopsi disetujui!");
                    location.reload();
                })
                .catch(err => console.error(err));
        });
    });

    rejectButtons.forEach(btn => {
        btn.addEventListener("click", function () {
            const id = this.getAttribute("data-id");
            fetch(`/dashboard/adoptions/${id}/reject`, { method: "POST" })
                .then(response => {
                    if (response.ok) {
                        return response.json();
                    } else {
                        throw new Error("Failed to reject adoption");
                    }
                })
                .then(data => {
                    alert("Adopsi ditolak!");
                    location.reload();
                })
                .catch(err => console.error(err));
        });
    });
});
