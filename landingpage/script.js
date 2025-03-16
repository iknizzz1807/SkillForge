// Chờ DOM load xong
document.addEventListener("DOMContentLoaded", () => {
  // 1. Skill Selector Logic
  const skillButtons = document.querySelectorAll(".skill-btn");
  const matchOutput = document.getElementById("match-output");
  let selectedSkills = [];

  skillButtons.forEach((button) => {
    button.addEventListener("click", () => {
      const skill = button.dataset.skill;
      if (selectedSkills.includes(skill)) {
        selectedSkills = selectedSkills.filter((s) => s !== skill);
        button.classList.remove("active");
      } else {
        selectedSkills.push(skill);
        button.classList.add("active");
      }

      // Cập nhật kết quả matching
      if (selectedSkills.length === 0) {
        matchOutput.textContent = "Select skills to see results!";
      } else {
        const projects = matchProjects(selectedSkills);
        matchOutput.textContent = projects;
        matchOutput.style.opacity = 0;
        setTimeout(() => {
          matchOutput.style.transition = "opacity 0.5s";
          matchOutput.style.opacity = 1;
        }, 100);
      }
    });
  });

  // Hàm giả lập matching projects (có thể thay bằng API thực tế sau)
  function matchProjects(skills) {
    const projectPool = {
      javascript: ["Build a Dynamic Web App", "Interactive Dashboard"],
      react: ["React E-commerce Site", "Social Media Frontend"],
      python: ["Data Analysis Tool", "Machine Learning Model"],
      nodejs: ["API Backend", "Real-time Chat App"],
    };

    let matched = [];
    skills.forEach((skill) => {
      if (projectPool[skill]) {
        matched = [...matched, ...projectPool[skill]];
      }
    });

    return matched.length > 0
      ? `We found "${matched[0]}" (Match: ${Math.min(
          95,
          70 + skills.length * 10
        )}%)`
      : "No projects match yet, try more skills!";
  }

  // 2. Testimonial Slider
  const testimonialItems = document.querySelectorAll(".testimonial-item");
  let currentTestimonial = 0;

  function showTestimonial(index) {
    testimonialItems.forEach((item, i) => {
      item.style.transform = `translateX(${(i - index) * 100}%)`;
      item.style.opacity = i === index ? 1 : 0;
    });
  }

  function nextTestimonial() {
    currentTestimonial = (currentTestimonial + 1) % testimonialItems.length;
    showTestimonial(currentTestimonial);
  }

  // Khởi động slider
  showTestimonial(currentTestimonial);
  setInterval(nextTestimonial, 5000); // Chuyển mỗi 5 giây

  // 3. Scroll Animations
  const sections = document.querySelectorAll("section");
  const observerOptions = {
    threshold: 0.2,
  };

  const sectionObserver = new IntersectionObserver((entries, observer) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        entry.target.classList.add("visible");
        observer.unobserve(entry.target);
      }
    });
  }, observerOptions);

  sections.forEach((section) => {
    section.classList.add("hidden");
    sectionObserver.observe(section);
  });

  // 5. Form Submission (Fake Feedback)
  const contactForm = document.querySelector(".contact-form");
  contactForm.addEventListener("submit", (e) => {
    e.preventDefault();
    const submitButton = contactForm.querySelector("button");
    submitButton.textContent = "Sending...";
    submitButton.disabled = true;

    setTimeout(() => {
      submitButton.textContent = "Message Sent!";
      submitButton.style.background = "#ffd700";
      setTimeout(() => {
        contactForm.reset();
        submitButton.textContent = "Send Message";
        submitButton.style.background = "#ff6f61";
        submitButton.disabled = false;
      }, 2000);
    }, 1000);
  });
});

// Thêm hiệu ứng khi scroll
window.addEventListener("scroll", () => {
  const header = document.querySelector(".header");
  if (window.scrollY > 50) {
    header.style.background = "linear-gradient(135deg, #5a3ecc, #e65a4d)";
    header.style.padding = "10px 0";
  } else {
    header.style.background = "linear-gradient(135deg, #6b48ff, #ff6f61)";
    header.style.padding = "20px 0";
  }
});
